# Exchanging an Auth Code for a Token

This Golang code sample demonstrates how to exchange an authentication code for a token using the sandbox app credentials obtained from Recipient Hub. It uses the Gin web framework and templates to render a basic HTML page with a button to start the consent flow process.

## Key Concepts
* Making a POST request to the identity provider to exchange an auth code for a token.
* Understanding the OAuth2 authorization code flow.
* Redirecting to a test bank, "Mikomo," to simulate the "consent flow" for account access permissions.
* Redirecting to Recipient Hub's token viewer page to receive the token after granting access.
* Using the token to make a GET request to the identity provider to fetch the token.

## Project Overview
The main.go module includes a single route, /, which displays a basic HTML page rendered from the ./templates/index.tmml template. This page includes an anchor link that points to a sandbox identity provider URL, utilizing your ClientId variable, so ensure it is correctly set to your client ID.

The go.mod file contains the application's Golang dependencies, while the run.sh file is a shell script that streamlines dependency fetching, application compilation, and binary execution into a single step for convenience.


## Running the Sample
### From Localhost
1. Initialize the ClientId and Secret variables in main.go with your sandbox app's client ID and secret.
2. Open a terminal in the sandbox.
3. Run the command chmod +x run.sh.
4. Run the command ./run.sh.
5. The application will be available at http://localhost:8123.
6. Follow the instructions in the browser once the application is opened.

### From CodeSandbox
1. Open the project in CodeSandbox.
2. Click the "Sign In" button in CodeSandbox.
3. Choose "Sign In with GitHub."
4. Authenticate using your GitHub credentials and grant CodeSandbox access to your repositories.
5. After signing in, click the dropdown next to the "Sign In" button in CodeSandbox and choose "Fork Project."
6. When prompted, search for this repository to fork: https://github.com/akoya-llc/go-exchange-auth-code-for-token.
7. Click the "Fork" button to fork the project.
8. The application will start running in CodeSandbox, after which you'll be prompted to open the preview via the "Preview: 8123" button.
9. Click the "Preview: 8123" button to open the application in the browser.
10. In the toolbar next to the preview URL bar, hover over the icons and look for the tooltip text "Open in a new tab," then click it.
11. The application will start running in a new browser window where you can interact with it.
