package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**

	假设有两个表：
	accounts 表（包含字段 id 主键， balance 账户余额）和
	transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	要求 ：
	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
	向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

**/

type Accounts struct {
	Id      uint `gorm:"primaryKey"`
	Balance float64
}
type Transactions struct {
	Id            uint `gorm:"primaryKey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test001?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}

	db.AutoMigrate(&Accounts{}, &Transactions{})
	// 创建AB账户
	// AccountA := Accounts{Id: 1, Balance: 200}
	// resultA := db.Create(&AccountA)
	// if resultA.Error != nil {
	// 	log.Println("A账户创建失败:", resultA.Error)
	// }
	// fmt.Printf("A账户创建成功，ID: %d\n", AccountA.Id)

	db.Model(&Accounts{}).Where("Id = ?", 1).Updates(Accounts{Balance: 1000})

	// AccountB := Accounts{Id: 2, Balance: 200}
	// resultB := db.Create(&AccountB)
	// if resultB.Error != nil {
	// 	log.Println("B账户创建失败:", resultA.Error)
	// }
	// fmt.Printf("B账户创建成功，ID: %d\n", AccountB.Id)

	// 转出账户ID
	fromID := uint(1)

	// 转入账户ID
	toID := uint(2)
	// 交易金额
	amount := 100.0

	db.Transaction(func(tx *gorm.DB) error {

		var fromAccount Accounts
		if err := tx.First(&fromAccount, fromID).Error; err != nil {

			return fmt.Errorf("查询转出账户失败: %v", err)
		}

		if fromAccount.Balance < amount {
			return fmt.Errorf("账户余额不足")
		}

		var toAccount Accounts

		if err := tx.First(&toAccount, toID).Error; err != nil {
			return fmt.Errorf("查询转ru账户失败: %v", err)
		}

		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-amount).Error; err != nil {
			return fmt.Errorf("扣除转出账户金额失败: %v", err)
		}

		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+amount).Error; err != nil {

			return fmt.Errorf("更新失败: %v", err)
		}

		Transaction1 := Transactions{FromAccountId: fromID, ToAccountId: toID, Amount: amount}

		if err := tx.Create(&Transaction1).Error; err != nil {
			return fmt.Errorf("交易记录保存失败: %v", err)
		}

		return nil

	})

}
