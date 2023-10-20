---
geometry: margin=25mm
---

# Node.js Scraper Challenge

Welcome to the Node.js Scraper Challenge! :tada:

In this challenge, your task is to write a web scraper in Node.js, either in TypeScript or JavaScript, to extract specific information from our testing site.
The website has a login page, and once you log in, it displays a table of orders.
Such orders might have a link to downloadable CSV file that contains details about the order, including items and their prices.

## Challenge Details

### Website Information

- **Website URL**: [https://node-challenge.fly.dev/](https://node-challenge.fly.dev/)

### Login Information

- **Username**: jdoe
- **Password**: jdoe

### Website Structure

1. **Login Page**: The initial page where you need to log in with the provided credentials.
2. **Orders Page**: After logging in, you will be directed to a page that displays a table of orders. Each row in the table represents an "Order."
3. **CSV Links**: Within each row of the order table, there is a link to download a CSV file that contains details about the order, including items and their prices.
                  **Beware that some orders may not have a CSV file available for download.**
4. **Rate Limiting**: The website has a rate limit.

## Challenge Objective

Your objective is to create a Node.js script in either TypeScript or JavaScript that accomplishes the following tasks:

1. Log in to the website using the provided username and password (username and password **CAN NOT** be hardcoded in the script).
2. Extract the Buyer Name and Total expenditure (amount of items * item price) for orders exclusively containing items priced above $50, from the CSV file.
   - If the order has not even a single item priced above $50, you should skip that order.
   - If an order does not have a CSV file available for download, you should skip that order.
3. You are free to use any Node.js libraries or frameworks to accomplish the challenge. However, the use of best suited libraries and frameworks will be considered as a plus.

## Submission

1. **GitHub Repository**: Create a **Private GitHub** repository for your project and share the repository URL with us.
2. Add someone from our team as a collaborator to your repository. (You should have been told who to add as a collaborator.)

## Evaluation

Your solution will be evaluated based on the following criteria:

1. **Functionality**: Does your code accomplish the tasks outlined in the challenge?

2. **Code Quality**: Is your code well-organized, readable, and well-documented?

3. **Error Handling**: Does your code handle potential errors gracefully, such as handling failed logins or missing CSV files?

4. **Efficiency**: How efficiently does your scraper operate? Does it minimize processing time and resource usage?

5. **Database Integration**: Does your code effectively interact with your chosen database for data storage?

6. **Adherence to Best Practices**: Does your code follow best practices for Node.js development?

Good luck with the Challenge! Feel free to reach out if you have any questions or need assistance during the challenge.

