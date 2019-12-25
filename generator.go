package main

import (
	"fmt"
	"log"
	"strings"
)

const SELECT = "SELECT * FROM %s WHERE %s"
const DELETE = "DELETE FROM %s WHERE %s"
const UPDATE = "UPDATE %s SET "
const INSERT = "INSERT INTO %s (%s) VALUES (%s)"

func generate(name string, fields []string) string {
	if name[:4] == "Find" {
		return selectGenerator("", name)
	} else if name[:6] == "Update" {
		return updateGenerator("test", name)
	} else if name[:6] == "Delete" {
		return deleteGenerator("", name)
	} else if name[:6] == "INSERT" {
		return insertGenerator("", fields)
	} else {
		log.Fatalf("Method name %s is not valid\n", name)
		return ""
	}
}

func selectGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(SELECT, tableName, strings.Join(paramsAndQuestions, " AND "))
}
func deleteGenerator(tableName, method string) string {
	queryParams := strings.Split(method, "By")[1]
	params := strings.Split(queryParams, "And")
	var paramsAndQuestions []string
	for p := range params {
		paramsAndQuestions = append(paramsAndQuestions, fmt.Sprintf("%s=?", params[p]))
	}
	return fmt.Sprintf(DELETE, tableName, strings.Join(paramsAndQuestions, "AND"))
}
func updateGenerator(tableName string, method string) string {
	selectParams := strings.Split(strings.Split(method, "BasedOn")[1], "And")
	updateParams := strings.Split(strings.Split(method, "BasedOn")[0][6:], "And")

	query := fmt.Sprintf(UPDATE, tableName)

	var updateParamslist []string
	for _, updateParam := range updateParams{
		updateParamslist = append(updateParamslist, fmt.Sprintf("%s = ?", updateParam))
	}

	query += strings.Join(updateParamslist, ", ") + " WHERE "

	var selectParamslist []string
	for _, selectParam := range selectParams{
		selectParamslist = append(selectParamslist, fmt.Sprintf("%s = ?", selectParam))
	}
	query += strings.Join(selectParamslist, " AND ")

	return query
}
func insertGenerator(tableName string, fields []string) string {
	var questions []string
	for range fields {
		questions = append(questions, "?")
	}
	return fmt.Sprintf(INSERT, tableName, strings.Join(fields, ", "),strings.Join(questions, ", "))
}