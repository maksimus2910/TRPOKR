const apiUrl = "http://localhost:8080/products";

async function fetchProducts() {
  const res = await fetch(apiUrl);
  const products = await res.json();
  const tbody = document.querySelector("#products-table tbody");
  tbody.innerHTML = "";

  products.forEach(({ id, category, name, quantity, price }) => {
    const tr = document.createElement("tr");

    tr.innerHTML = `
      <td>${category}</td>
      <td>${name}</td>
      <td>${quantity}</td>
      <td>${price.toFixed(2)}</td>
      <td><button data-id="${id}">Удалить</button></td>
    `;

    tr.querySelector("button").addEventListener("click", async () => {
      await fetch(`${apiUrl}/${id}`, { method: "DELETE" });
      fetchProducts();
    });

    tbody.appendChild(tr);
  });
}

document.getElementById("add-product-form").addEventListener("submit", async (e) => {
  e.preventDefault();

  const newProduct = {
    category: e.target.category.value,
    name: e.target.name.value,
    quantity: parseInt(e.target.quantity.value),
    price: parseFloat(e.target.price.value),
  };

  await fetch(apiUrl, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(newProduct),
  });

  e.target.reset();
  fetchProducts();
});

fetchProducts();
