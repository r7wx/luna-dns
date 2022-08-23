/*
MIT License
Copyright (c) 2022 r7wx
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package engine

type node struct {
	key       string
	ip        string
	childrens []node
}

type domainTree struct {
	tlds []node
}

func (t *domainTree) insert(entry *entry) {
	foundTLD, _ := searchNode(t.tlds, entry.tld)
	if foundTLD == nil {
		newNode := node{
			key: entry.tld,
		}
		if entry.tld == "*" {
			newNode.ip = entry.ip
		}
		t.tlds, foundTLD = t.insertNode(t.tlds, newNode)
	}

	current := foundTLD
	for index, subdomain := range entry.subdomains {
		foundNode, _ := searchNode(current.childrens, subdomain)
		if foundNode == nil {
			newNode := node{
				key: subdomain,
			}
			if index == len(entry.subdomains)-1 {
				newNode.ip = entry.ip
			}
			current.childrens, foundNode = t.insertNode(current.childrens, newNode)
		}
		current = foundNode
	}
}

func (t *domainTree) search(entry *entry) string {
	foundTLD, wildcard := searchNode(t.tlds, entry.tld)
	if wildcard {
		return foundTLD.ip
	}

	if foundTLD != nil {
		current := foundTLD
		for index, subdomain := range entry.subdomains {
			foundNode, wildcard := searchNode(current.childrens, subdomain)
			switch {
			case wildcard:
				return foundNode.ip
			case foundNode == nil:
				return ""
			case index == len(entry.subdomains)-1:
				return foundNode.ip
			}
			current = foundNode
		}
	}

	return ""
}

func (t *domainTree) insertNode(nodes []node, newNode node) ([]node, *node) {
	if newNode.key == "*" {
		output := []node{newNode}
		outnode, _ := searchNode(output, newNode.key)
		return output, outnode
	}

	position := 0
	for _, node := range nodes {
		if node.key > newNode.key {
			break
		}
		position++
	}

	left := nodes[:position]
	right := nodes[position:]

	output := []node{}
	output = append(output, left...)
	output = append(output, newNode)
	output = append(output, right...)

	outnode, _ := searchNode(output, newNode.key)
	return output, outnode
}

func searchNode(nodes []node, key string) (*node, bool) {
	high := len(nodes) - 1
	low := 0

	for low <= high {
		pivot := (low + high) / 2
		if nodes[pivot].key == "*" {
			return &nodes[pivot], true
		}
		if nodes[pivot].key < key {
			low = pivot + 1
		} else {
			high = pivot - 1
		}
	}

	if low == len(nodes) || nodes[low].key != key {
		return nil, false
	}

	return &nodes[low], false
}
