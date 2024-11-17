// Sample cart data (this would be loaded from a backend or database in a real application)
let api = 'localhost:8080';
let cart = [
    { id: 1, name: "Product 1", price: 10.00, quantity: 1 },
    { id: 2, name: "Product 2", price: 15.00, quantity: 1 }
];

// Function to render cart items on page load
function renderCart() {
    const cartItems = document.getElementById("cart-items");
    cartItems.innerHTML = ''; // Clear existing items

    cart.forEach(item => {
        const li = document.createElement("li");
        li.innerHTML = `
            ${item.name} - $${item.price.toFixed(2)} x 
            <input type="number" min="1" value="${item.quantity}" onchange="updateQuantity(${item.id}, this.value)">
            <button onclick="removeItem(${item.id})">Remove</button>
        `;
        cartItems.appendChild(li);
    });

    updateCartSummary();
}

// Update item quantity
function updateQuantity(itemId, newQuantity) {
    const item = cart.find(product => product.id === itemId);
    if (item) {
        item.quantity = parseInt(newQuantity);
        updateCartSummary();
    }
}

// Remove an item from the cart
function removeItem(itemId) {
    cart = cart.filter(product => product.id !== itemId);
    renderCart();
}

// Update cart summary and total price
function updateCartSummary() {
    const total = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);
    document.getElementById("cart-count").textContent = cart.length;
    document.getElementById("cart-total").textContent = total.toFixed(2);
}

// Checkout function
function checkout() {
    alert("Proceeding to checkout with a total of $" + 
        cart.reduce((sum, item) => sum + item.price * item.quantity, 0).toFixed(2));
    // Send cart data to api
    fetch(api+'/orders/new', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(cart)
    }).then(response => {
        // Handle api response
        if (!response.ok) {
            alert("Error placing order.");
            console.log(response);
        }
        alert("Order placed successfully.");
    });
    cart = []; // Clear the cart after checkout
    renderCart();
}

// Initialize cart on page load
window.onload = renderCart;
