let api = 'https://www.tesc.farm/api';
let cart = [
    { id: 1, name: "Product 1", price: 10.00, quantity: 1 },
    { id: 2, name: "Product 2", price: 15.00, quantity: 1 }
];

// Render cart items
function renderCart() {
    const cartItems = document.getElementById("cart-items");
    if (!cartItems) return;

    cartItems.innerHTML = ''; // Clear existing items
    cart.forEach(item => {
        const li = document.createElement("li");
        li.innerHTML = `
            ${item.name} - $${item.price.toFixed(2)} x 
            <input type="number" min="1" value="${item.quantity}" 
                   onchange="updateQuantity(${item.id}, this.value)">
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
        item.quantity = Math.max(parseInt(newQuantity) || 1, 1);
        updateCartSummary();
        renderCart();
    }
}

// Remove an item from the cart
function removeItem(itemId) {
    cart = cart.filter(product => product.id !== itemId);
    renderCart();
}

// Update cart summary
function updateCartSummary() {
    const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);
    const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);
    document.getElementById("cart-count").textContent = totalItems;
    document.getElementById("cart-total").textContent = totalPrice.toFixed(2);
}

// Checkout function
function checkout() {
    // Retrieve cart total
    const total = cart.reduce((sum, item) => sum + item.price * item.quantity, 0).toFixed(2);

    if (cart.length === 0) {
        alert("Your cart is empty!");
        return;
    }

    // Capture user information
    const name = document.getElementById('name').value;
    const email = document.getElementById('email').value;
    const address = document.getElementById('address').value;
    const city = document.getElementById('city').value;
    const state = document.getElementById('state').value;
    const zip = document.getElementById('zip').value;
    const card = document.getElementById('card').value;
    const expiry = document.getElementById('expiry').value;
    const cvv = document.getElementById('cvv').value;

    // Basic validation
    if (!name || !email || !address || !city || !state || !zip || !card || !expiry || !cvv) {
        alert("Please fill out all required fields.");
        return;
    }

    if (!confirm(`Proceed to checkout with a total of $${total}?`)) {
        return; // Cancel checkout if user declines
    }

    // Prepare payload with user info and cart data
    const payload = {
        customer: {
            name,
            email,
            address: {
                street: address,
                city,
                state,
                zip
            }
        },
        payment: {
            cardNumber: card,
            expiry,
            cvv
        },
        items: cart.map(item => ({
            product: item.id,
            quantity: item.quantity
        }))
    };

    // Send data to API
    fetch(`${api}/orders/new`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error("Error placing order. Please try again.");
        }
        return response.json();
    })
    .then(data => {
        alert(`Order placed successfully! Order ID: ${data.orderId}`);
        cart = []; // Clear the cart on success
        renderCart();
    })
    .catch(error => {
        console.error('Error:', error);
        alert(error.message);
    });
}

// Initialize cart on page load
window.onload = renderCart;
