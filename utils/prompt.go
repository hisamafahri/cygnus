package utils

import (
    "fmt"

    "github.com/AlecAivazis/survey/v2"
    )

func PromptText(question *string) (string, error) {
    answers := struct {
        Input          string
    }{}

    qs := []*survey.Question{
        {
            Name:     "input",
            Prompt:   &survey.Input{Message: *question},
            Validate: survey.Required,
            Transform: survey.Title,
        },
    }

    // perform the questions
    err := survey.Ask(qs, &answers)
    if err != nil {
        fmt.Printf(" error: %s", err)
        return "", err
    }

    return answers.Input, nil
}
