let cart = [];
let cartTotal = 0;

function addToCart(productName, productPrice) {
    const product = { name: productName, price: productPrice };
    cart.push(product);
    cartTotal += productPrice;
    updateCartSummary();
    showCart();
}

function updateCartSummary() {
    document.getElementById("cart-count").textContent = cart.length;
    document.getElementById("cart-total").textContent = cartTotal.toFixed(2);
    updateCartItems();
}

function updateCartItems() {
    const cartItems = document.getElementById("cart-items");
    cartItems.innerHTML = '';
    cart.forEach(item => {
        const li = document.createElement("li");
        li.textContent = `${item.name} - $${item.price.toFixed(2)}`;
        cartItems.appendChild(li);
    });
}

function showCart() {
    document.getElementById("cart-modal").classList.remove("hidden");
}

function closeCart() {
    document.getElementById("cart-modal").classList.add("hidden");
}

function checkout() {
    alert("Proceeding to checkout with a total of $" + cartTotal.toFixed(2));
    cart = [];
    cartTotal = 0;
    updateCartSummary();
    closeCart();
}