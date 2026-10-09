package servicebus

import (
	"context"
	"encoding/json"
	"fmt"

	"notification-agent/internal/models"
	"notification-agent/internal/repository"
	"notification-agent/internal/service"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func ReceiveMessage(
	connectionString string,
	queueName string,
) error {

	client, err := azservicebus.NewClientFromConnectionString(
		connectionString,
		nil,
	)

	if err != nil {
		return err
	}

	receiver, err := client.NewReceiverForQueue(
		queueName,
		nil,
	)

	if err != nil {
		return err
	}

	msgs, err := receiver.ReceiveMessages(
		context.Background(),
		1,
		nil,
	)

	if err != nil {
		return err
	}

	for _, msg := range msgs {

		fmt.Println("=================================")
		fmt.Println("Message Received From Queue")

		fmt.Println(
			"Azure Message ID:",
			msg.MessageID,
		)

		var notification models.Notification

		err := json.Unmarshal(
			msg.Body,
			&notification,
		)

		fmt.Printf("%+v\n", notification)

		if err != nil {

			fmt.Println("JSON Parse Failed")

			receiver.DeadLetterMessage(
				context.Background(),
				msg,
				nil,
			)

			continue
		}

		err = service.ProcessNotification(
			notification,
		)

		if err != nil {

			fmt.Println(err)

			receiver.DeadLetterMessage(
				context.Background(),
				msg,
				nil,
			)

			continue
		}

		repository.SaveMessageLog(
			notification.NotificationID,
			msg.MessageID,
			"RECEIVED",
		)

		fmt.Println(
			"Message Log Saved Successfully",
		)

		fmt.Println(
			"Azure Message ID:",
			msg.MessageID,
		)

		fmt.Println("=================================")

		err = receiver.CompleteMessage(
			context.Background(),
			msg,
			nil,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
