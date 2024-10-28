# 💸 Budget Tracker CLI

A simple and efficient command-line application for tracking your income and expenses, developed in Go. This CLI app allows you to add transactions, display all transactions, show total income and expenses, and save transactions to a CSV file for future reference.

## ✨ Features

- 📥 **Add Transaction**: Record income and expense transactions easily.
- 📋 **Display Transactions**: View all transactions along with details.
- 💰 **Show Total Income**: Calculate and display the total income.
- 🧾 **Show Total Expenses**: Calculate and display the total expenses.
- 💾 **Save to CSV**: Export all transactions to a CSV file.

## ⚙️ Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/YashSaini99/Budget_Tracker.git
   cd Budget_Tracker
   ```

2. **Build the application:**
   ```bash
   go build -o main
   ```

## 📄 CSV Format

The CSV file will contain the following columns:
- 📅 **Date** - The date of the transaction.
- 📈 **Type** - The type of transaction (`income` or `expense`).
- 💵 **Amount** - The amount of the transaction.
- 📝 **Description** - The description of the transaction.


## Screenshot

![App Screenshot](https://github.com/YashSaini99/Buget_Tracker/blob/main/CLI.png)

## 🛠 Requirements

- Go 1.16 or higher

## 📜 License

This project is licensed under the MIT License.
