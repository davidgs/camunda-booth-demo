# Camunda Booth Demo

This is a repo for a Camunda Platform 8 demo for trade show booths.

## Idea

Given that Niall is a Quizmaster, we thought maybe a fun demo would be a quiz app.

The idea is that we will use Camunda Platform 8 to retrieve topics and questions from 3 open-source trivia APIs.

  * [Open Trivia API](https://opentdb.com/)
  * [The Trivia API](https://the-trivia-api.com/)
  * [Jeopardy API](https://github.com/sottenad/jService)

I have added a preliminary BPMN diagram to the repo, along with an initial Go client to fetch categories from the-trivia-api.com.

I have a POC of being able to call (and be called from) Zeebe APIs from a React.JS app. I will add that code here shortly.