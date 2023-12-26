package main

import (
	"context"
	"github.com/jopoleon/ClassroomBot/internal/clients/github_classroom"
	"github.com/jopoleon/ClassroomBot/internal/server/telegram_bot"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
)

func main() {

	botC, err := telegram_bot.NewClassroomTrackerBot(
		"6763994595:AAG2biFynwlJaaMUiFQN7wOYWV9g_Sebd1k",
		"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",
	)
	if err != nil {
		logrus.Error("ListClassrooms err ", err)
		return
	}
	botC.Start()

	return

	//emiloserdov93
	accessToken := "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"

	classroomClient := github_classroom.NewClient(context.Background(), accessToken)
	crs, err := classroomClient.ListClassrooms()
	if err != nil {
		logrus.Error("ListClassrooms err ", err)
		return
	}

	pp.Println(crs)

	assignments, err := classroomClient.GetClassroomAssignments(crs[0])
	if err != nil {
		logrus.Error("ListClassrooms err ", err)
		return
	}

	//pp.Println(assignments)

	for _, a := range assignments {
		if a.Id == 533251 {
			ass, err := classroomClient.GetAssignmentInfo(a.Id)
			if err != nil {
				logrus.Error("GetAssignmentInfo err ", err)
				return
			}
			pp.Println(ass)
			aFull, err := classroomClient.GetAcceptedAssignments(a.Id)
			if err != nil {
				logrus.Error("GetAcceptedAssignments err ", err)
				return
			}

			pp.Println(aFull)
		}
	}
	//533251

}
