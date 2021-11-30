package main

import (
	"fmt"
	"sort"
)

type DSU struct {
	representative []int
	size           []int
}

func (d *DSU) DSU(sz int) {
	d.representative = make([]int, sz)
	d.size = make([]int, sz)

	for i := 0; i < sz; i++ {
		// Initially each group is its own representative
		d.representative[i] = i
		// Intialize the size of all groups to 1
		d.size[i] = 1
	}
}

// Finds the representative of group x
func (d *DSU) findRepresentative(x int) int {
	if x == d.representative[x] {
		return x
	}

	d.representative[x] = d.findRepresentative(d.representative[x])
	return d.representative[x]
}

// Unite the group that contains "a" with the group that contains "b"
func (d *DSU) unionBySize(a int, b int) {
	representativeA := d.findRepresentative(a)
	representativeB := d.findRepresentative(b)

	// If nodes a and b already belong to the same group, do nothing.
	if representativeA == representativeB {
		return
	}

	// Union by size: point the representative of the smaller
	// group to the representative of the larger group.
	if d.size[representativeA] >= d.size[representativeB] {
		d.size[representativeA] += d.size[representativeB]
		d.representative[representativeB] = representativeA
	} else {
		d.size[representativeB] += d.size[representativeA]
		d.representative[representativeA] = representativeB
	}
}

func accountsMerge(accounts [][]string) [][]string {
	accountListSize := len(accounts)
	d := DSU{}
	d.DSU(accountListSize)

	// Maps email to their component index
	emailGroup := make(map[string]int)

	for i := 0; i < accountListSize; i++ {
		accountSize := len(accounts[i])

		for j := 1; j < accountSize; j++ {
			email := accounts[i][j]

			// If this is the first time seeing this email then
			// assign component group as the account index
			if _, ok := emailGroup[email]; ok {
				// If we have seen this email before then union this
				// group with the previous group of the email
				d.unionBySize(i, emailGroup[email])
			} else {
				emailGroup[email] = i
			}
		}
	}

	// Store emails corresponding to the component's representative
	components := make(map[int][]string)
	for email := range emailGroup {
		group := emailGroup[email]
		groupRep := d.findRepresentative(group)

		if _, ok := components[groupRep]; !ok {
			components[groupRep] = []string{}
		}

		components[groupRep] = append(components[groupRep], email)
	}

	// Sort the components and add the account name
	mergedAccounts := [][]string{}
	for group := range components {
		component := components[group]
		sort.Strings(component)
		component = append([]string{accounts[group][0]}, component...)
		mergedAccounts = append(mergedAccounts, component)
	}

	return mergedAccounts
}

func main() {

	fmt.Println(combine(20, 10))
}
