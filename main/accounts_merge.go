package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/accounts-merge/

Given a list of accounts where each element accounts[i] is a list of strings,
where the first element accounts[i][0] is a name, and the rest of the elements are emails representing emails of the account.

Now, we would like to merge these accounts.
Two accounts definitely belong to the same person if there is some common email to both accounts.
Note that even if two accounts have the same name, they may belong to different people as people could have the same name.
A person can have any number of accounts initially, but all of their accounts definitely have the same name.

After merging the accounts, return the accounts in the following format:
the first element of each account is the name, and the rest of the elements are emails in sorted order.
The accounts themselves can be returned in any order.



Example 1:

Input: accounts = [
	["John","johnsmith@mail.com","john_newyork@mail.com"],
	["John","johnsmith@mail.com","john00@mail.com"],
	["Mary","mary@mail.com"],
	["John","johnnybravo@mail.com"]
]

Output: [
	["John","john00@mail.com","john_newyork@mail.com","johnsmith@mail.com"],
	["Mary","mary@mail.com"],
	["John","johnnybravo@mail.com"]
]
Explanation:
The first and third John's are the same person as they have the common email "johnsmith@mail.com".
The second John and Mary are different people as none of their email addresses are used by other accounts.
We could return these lists in any order, for example the answer [['Mary', 'mary@mail.com'], ['John', 'johnnybravo@mail.com'],
['John', 'john00@mail.com', 'john_newyork@mail.com', 'johnsmith@mail.com']] would still be accepted.

Example 2:
Input: accounts = [
	["Gabe","Gabe0@m.co","Gabe3@m.co","Gabe1@m.co"],
	["Kevin","Kevin3@m.co","Kevin5@m.co","Kevin0@m.co"],
	["Ethan","Ethan5@m.co","Ethan4@m.co","Ethan0@m.co"],
	["Hanzo","Hanzo3@m.co","Hanzo1@m.co","Hanzo0@m.co"],
	["Fern","Fern5@m.co","Fern1@m.co","Fern0@m.co"]
]
Output: [
	["Ethan","Ethan0@m.co","Ethan4@m.co","Ethan5@m.co"],
	["Gabe","Gabe0@m.co","Gabe1@m.co","Gabe3@m.co"],
	["Hanzo","Hanzo0@m.co","Hanzo1@m.co","Hanzo3@m.co"],
	["Kevin","Kevin0@m.co","Kevin3@m.co","Kevin5@m.co"],
	["Fern","Fern0@m.co","Fern1@m.co","Fern5@m.co"]
]


*/

func main() {
	fmt.Println("initializing test data...")

	accounts := [][]string{
		{"John","johnsmith@mail.com","john_newyork@mail.com"},
		{"John","johnsmith@mail.com","john00@mail.com"},
		{"Mary","mary@mail.com"},
		{"John","johnnybravo@mail.com"},
	}
	fmt.Println("Answer is:", accountsMerge(accounts))

	accounts2 := [][]string{
		{"David", "David0@m.co", "David1@m.co"},
		{"David", "David3@m.co", "David4@m.co"},
		{"David","David4@m.co","David5@m.co"},
		{"David","David2@m.co","David3@m.co"},
		{"David","David1@m.co","David2@m.co"},
	}

	fmt.Println("Answer is:", accountsMerge(accounts2))
}

func accountsMerge(accounts [][]string) [][]string {
	accountMap := make(map[string]*[]string)
	var result []*[]string

	for _, account := range accounts {
		fmt.Println("dealing with account:", account)
		var existingAccount *[]string = nil
		for i:=1; i<len(account); i++ {
			if accountMap[account[i]] != nil {
				existingAccount = accountMap[account[i]]
				break
			}
		}
		if existingAccount == nil {
			newAccount := &[]string{account[0]}
			for i:=1; i<len(account); i++ {
				//input list can have duplicates within a given account too
				if accountMap[account[i]] == nil {
					*newAccount = append(*newAccount, account[i])
					accountMap[account[i]] = newAccount
				}
			}
			result = append(result, newAccount)
		} else {
			for i:=1; i<len(account); i++ {
				if accountMap[account[i]] == nil {
					accountMap[account[i]] = existingAccount
					*existingAccount = append(*existingAccount, account[i])
				}
			}
		}
	}
	//prepare, sort and return result
	var sortedList [][]string
	for _, r := range result {
		emails := (*r)[1:]
		sort.Slice(emails, func(i, j int) bool {
			return emails[i] < emails[j]
		})
		sortedList = append(sortedList, *r)
	}
	return sortedList
}

func accountsMergeUnionFind(accounts [][]string) [][]string {
	//TODO
	return [][]string{{}}
}