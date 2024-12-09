package sender

/*
// Тест для структуры Sender
func TestSender(t *testing.T) {
	queue := &Queue{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockMessageStorager(ctrl)

	sender := NewSender(1, queue, storage, 587, "smtp.example.com", "test@example.com", "password")

	// Проверка, что Sender создан успешно
	if sender.ID != 1 {
		t.Errorf("expected ID to be 1, got %d", sender.ID)
	}

	msg := &core.Message{
		Destination: "recipient@example.com",
		Login:       "Test Subject",
		Text:        "This is a test email body.",
		Attachment:  "test_attachment.png",
	}
	queue.Push(msg)

	exit := make(chan struct{})
	go sender.StartSending(exit)

	time.Sleep(1 * time.Second) // Даем немного времени на отправку

	close(exit) // Закрываем канал exit
	// Здесь вы можете добавить дополнительные проверки,
	// например, проверить, что сообщение было добавлено в storage или было отправлено
}
*/
