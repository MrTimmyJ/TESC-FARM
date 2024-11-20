let api = 'https://www.tesc.farm/api';
let cart = [
    { id: 1, name: "Product 1", price: 10.00, quantity: 1 },
    { id: 2, name: "Product 2", price: 15.00, quantity: 1 },
    { id: 3, name: "Product 2", price: 15.00, quantity: 1 },
    { id: 4, name: "Product 2", price: 15.00, quantity: 1 },
    { id: 5, name: "Product 2", price: 15.00, quantity: 1 }
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

function renderProduce() {
    let grids = document.getElementsByClassName("product-grid");
    if (grids.length < 1) {
        return;
    }

    let grid = grids[0];

    fetch(`${api}/products`, {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
        if (!response.ok) {
            throw new Error("Error finding products. Please try again.");
        }
        return response.json();
    }).then(data => {
        grid.innerHTML = "";
        for (const product of data.Products) {
            const card = document.createElement("div");
            card.classList.add("product-card");

            const prod_img = document.createElement("img");
            prod_img.src = "static/img/products/" + product.image;
            prod_img.alt = product.name;
            card.appendChild(prod_img);

            const card_details = document.createElement("div");
            card_details.classList.add("product-details");
            

            const prod_name = document.createElement("h3");
            prod_name.classList.add("product-name");
            prod_name.innerText = product.name;
            card_details.appendChild(prod_name);
            const prod_unit = document.createElement("h3");
            prod_unit.classList.add("product-unit");
            prod_unit.innerText = product.unit;
            card_details.appendChild(prod_unit);
            card.appendChild(card_details);

            const card_checkout = document.createElement("div");
            card_checkout.classList.add("product-card-checkout");

            const prod_price = document.createElement("p");
            prod_price.classList.add("price");
            prod_price.innerText = "$" + product.price.toFixed(2);
            card_checkout.appendChild(prod_price);

            const prod_id = document.createElement("input");
            prod_id.type = "hidden";
            prod_id.name = "prod_id";
            prod_id.value = product.id;
            card_checkout.appendChild(prod_id);

            const prod_button = document.createElement("button");
            prod_button.classList.add("add-to-cart-button");
            prod_button.textContent = "Add to Cart";
            card_checkout.appendChild(prod_button);
            card.appendChild(card_checkout);

            grid.appendChild(card);

            prod_button.addEventListener("click", addToCart);
        }
    }).catch(error => {
        console.error('Error:', error);
        alert(error.message);
    });
}

// Add item to cart using 'e' as a Button Click Event
function addToCart(e) {
    let id = e.target.parentElement.querySelector("input[name='prod_id']").value;

    let tempCart = localStorage.getItem("cart");
    if (tempCart == null) {
        tempCart = [];
    } else {
        for (const item of tempCart) {
            if (item.id == id) {
                item.quantity++;
                localStorage.setItem("cart", tempCart);
                return;
            }
        }
    }

    let name = e.target.parentElement.parentElement.querySelector(".product-name").innerText;
    let price = e.target.parentElement.querySelector(".price").innerText;

    tempCart.append({"id": id, "name": name, "price": price, "quantity": 1});
    localStorage.setItem("cart", tempCart);
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
    }).then(response => {
        if (!response.ok) {
            throw new Error("Error placing order. Please try again.");
        }
        return response.json();
    }).then(data => {
        alert(`Order placed successfully! Order ID: ${data.orderId}`);
        cart = []; // Clear the cart on success
        renderCart();
    }).catch(error => {
        console.error('Error:', error);
        alert(error.message);
    });
}

// Initialize cart on page load
window.onload = renderCart;
window.addEventListener("load", function(){
    renderCart();

    renderProduce();
});
