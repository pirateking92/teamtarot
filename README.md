<h1>Team Tarot</h1>

Welcome to the README for Team Tarot, you're new favourite fortune teller

This repo contains the codebase for the Tarot Card Reading 
Application made by Team Tarot (Alannah Lawlor, Josh Day, Fara Ifaturoti, Oli Kelly,
Giandomenico Loparco and Matt Doyle).

## Table of Contents

- [Description](#description)
- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Project Structure](#project-structure)
- [Acknowledgements](#acknowledgements)
- [Contact](#contact)

## Description
Team Tarot is a tarot card reading application that, when the user is ready, will deal 3 random cards from a 78 card deck of tarot cards, provide an explanation for each, and then give a reading. Each card will appear reversed or upright and their meanings will be displayed accordingly. The type of card whether major or minor will also be printed, each of the three tarot cards printed will provide information about the past, present and the future.

## Installation
To run Team Tarot locally, you'll need to install:
- [Golang](https://go.dev/) 
- [Gin Framework](https://go.dev/doc/modules/) 
- [SvelteKit](https://github.com/sveltejs/kit) 
- [Node](https://nodejs.org/en)

## Usage
1. Clone the repository `git clone https://github.com/pirateking92/teamtarot` 

2. Go to the project directory `cd teamtarot`

3. Install dependencies `npm install ` 

4. Run the backend server `go run main.go`

5. Run the frontend server `npm run dev`

6. Navigate to `http://localhost:5173` 

## Features
- Users provide their name and a user story.
- The application generates 3 random tarot cards facing down.
- Upon clicking, the application reveals 3 random tarot cards along with their names, types, and descriptions, based on whether the card is displayed upright or reversed.
- The generated tarot cards and the user story are used to generate a personalized tarot card reading.

## Technologies Used
- GoLang is used for the backend with a Gin framework to handle HTTP requests. 
- SvelteKit is used for the frontend, requiring that node.js be installed.

## Project Structure 
Frontend 
- .svelte-kit: This directory contains configuration files and settings specific to the SvelteKit framework. It includes files related to routing, server setup, and other SvelteKit-specific configurations.
- node_modules: This directory holds the dependencies required by the frontend. These dependencies are installed via npm and include libraries and packages.
- src: Where the bulk of the frontend code is. It contains the source code, including Svelte components, JavaScript modules, CSS files, and other assets.
- static: This directory stores static assets.

Backend 
- controllers: This directory hold the logic behing incomming HTTP requests, interacts with models and generates responses 
- env: Holds functionality for loading environment variables
- models: Defines the properties of the tarot cards allowing for manipulaton of the tarot cards within the application 
- routes: This contains files that define API endpoints and map them to corresponding controllers 
- services: Holds functions for fetching tarot card data from an external API and interpreting them using the OpenAI API 

## Diagrams

...

## Acknowledgements



## Contact


