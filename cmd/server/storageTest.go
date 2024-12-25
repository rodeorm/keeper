package main

import (
	"context"
	"fmt"

	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/core"
)

func TestStorage(ctx context.Context, srv *cfg.Server) {

	usr1 := &core.User{Login: "user1", Password: "12345", Email: "ilyin-a-l@yandex.ru"}

	srv.UserStorager.RegUser(ctx, usr1)
	b := &core.Binary{Value: []byte("something"), Meta: "First item"}
	card := &core.Card{CardNumber: "2132 3423 4233 1234", OwnerName: "ALEXANDER ILYIN", Meta: "Tis"}
	text := &core.Text{Value: "Something interesting", Meta: "Second item"}
	couple := &core.Couple{Source: "gosuslugi", Login: "yandex", Password: "strongpassword"}

	srv.BinaryStorager.AddBinaryByUser(ctx, b, usr1)
	srv.CardStorager.AddCardByUser(ctx, card, usr1)
	srv.CoupleStorager.AddCoupleByUser(ctx, couple, usr1)
	srv.TextStorager.AddTextByUser(ctx, text, usr1)

	cards, _ := srv.CardStorager.SelectAllCardsByUser(ctx, usr1)
	for i, v := range cards {
		fmt.Println(i, v)
	}
	bins, _ := srv.BinaryStorager.SelectAllBinariesByUser(ctx, usr1)
	for i, v := range bins {
		fmt.Println(i, string(v.Value), v.Meta)
	}
	couples, _ := srv.CoupleStorager.SelectAllCouplesByUser(ctx, usr1)
	for i, v := range couples {
		fmt.Println(i, v)
	}
	texts, _ := srv.TextStorager.SelectAllTextsByUser(ctx, usr1)
	for i, v := range texts {
		fmt.Println(i, v)
	}

}
