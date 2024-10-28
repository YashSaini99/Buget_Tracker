package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Transaction struct{
	ID int
	Amount float64
	Category string
	Date time.Time
	Type string
}

type BugetTracker struct {
	transactions []Transaction
	nextID int
}

type FinancialRecord interface{
	GetAmount() float64
	GetType() string
}

func (t Transaction) GetAmount() float64{
	return t.Amount
}

func (t Transaction) GetType() string{
	return t.Type
}

func (bt *BugetTracker) AddTransaction(amount float64,category ,tType string){
	newTransaction := Transaction{
		ID: bt.nextID,
        Amount: amount,
        Category: category,
        Date: time.Now(),
        Type: tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)	
	bt.nextID++
}

func (bt BugetTracker) DisplayTransactions(){
	fmt.Println("ID\tAmount\tcategory\tDate\tType")
	for _, t := range bt.transactions{
        fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n", t.ID, t.Amount, t.Category, t.Date.Format("2006-01-02"), t.Type)
    }
}

func (bt BugetTracker) TotalIncome(tType string) float64{
	var total float64
    for _, t := range bt.transactions{
        if t.Type == tType{
            total += t.Amount
        }
    }
    return total
}

func (bt BugetTracker) SaveToCSV(filename string) error{
	file,err := os.Create(filename)
	if err!= nil{
        return err
    }
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID","Amount","Category","Date","Type"})

	for _, t := range bt.transactions{
        record := []string{strconv.Itoa(t.ID), fmt.Sprintf("%.2f", t.Amount), t.Category, t.Date.Format("2006-01-02"), t.Type}
        writer.Write(record)
    }
	fmt.Println("Transactions SAVED TO: ", filename)

	return nil
}

func main(){
	bt := BugetTracker{}
	for {
		fmt.Println("\n---Personal Buget Tracker---")
        fmt.Println("1. Add Transaction")
        fmt.Println("2. Display Transactions")
        fmt.Println("3. Show Total Income")
        fmt.Println("4  Show Total Expenses")
        fmt.Println("5. Save to CSV file")
		fmt.Println("6. Exit")
		fmt.Println("Choose an Option: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            var amount float64
            var category, tType string
            fmt.Print("Enter Amount: ")
            fmt.Scan(&amount)
            fmt.Print("Enter Category: ")
            fmt.Scan(&category)
            fmt.Print("Enter Type (Income/Expense): ")
            fmt.Scan(&tType)

            bt.AddTransaction(amount, category, tType)
			fmt.Print("Transaction Added!")
        case 2:
            bt.DisplayTransactions()
        case 3:
			fmt.Printf("Total Income: %.2f\n", bt.TotalIncome("income"))
		case 4:
			fmt.Printf("Total Expenses: %.2f\n", bt.TotalIncome("expense"))
        case 5:
			fmt.Printf("Enter File name (e.g. transactions	.csv): ")
			var filename string
			fmt.Scanln(&filename)
			if err := bt.SaveToCSV(filename); err != nil {
				fmt.Println("Error saving to CSV:", err)
			}
		case 6:
			fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid Option!")
        }
        fmt.Println("\n--- --- --- --- --- --- --- --- --- --- --- --- --- --- ---")
		}
	}

