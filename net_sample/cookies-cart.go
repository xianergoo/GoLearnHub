package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func splt(s, seq string) []string {
	return strings.Split(s, seq)
}

func parseCart(cartStr string) map[string]int {
	cart := make(map[string]int)
	for _, item := range splt(cartStr, ",") {
		if item == "" {
			continue
		}
		parts := splt(item, ":")
		if len(parts) != 2 {
			continue
		}

		id := parts[0]
		quantity, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		cart[id] += quantity
	}
	return cart
}

func join(s []string, seq string) string {
	return strings.Join(s, seq)
}

func formatCart(cart map[string]int) string {
	items := make([]string, 0, len(cart))
	for itemID, quantity := range cart {
		items = append(items, itemID+":"+strconv.Itoa(quantity))
	}
	return join(items, ",")

}

func add_to_cart(w http.ResponseWriter, r *http.Request) {
	itemID := r.FormValue("item_id")
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	cookie, err := r.Cookie("cart")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "cart",
			Value: "",
		}
	}

	//add to cart
	cart := parseCart(cookie.Value)
	cart[itemID] += quantity
	cookie.Value = formatCart(cart)
	http.SetCookie(w, cookie)

	fmt.Fprint(w, "item added to cart.")

}

func view_cart(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("cart")
	if err == http.ErrNoCookie {
		fmt.Fprint(w, "you cart is empty.")
		return
	}

	cart := parseCart(cookie.Value)
	if len(cart) == 0 {
		fmt.Fprint(w, "your cart is empty")
		return
	}

	fmt.Fprintln(w, "Your cart: ")
	for itemID, quantity := range cart {
		fmt.Fprintf(w, "- item %s: %d\n", itemID, quantity)
	}

}

func main() {
	http.HandleFunc("/add-to-cart", add_to_cart)
	http.HandleFunc("/view-cart", view_cart)
	http.ListenAndServe(":8080", nil)

}
