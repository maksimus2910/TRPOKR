CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  category VARCHAR(100),
  name VARCHAR(255),
  quantity INT,
  price NUMERIC(10,2)
);


CREATE OR REPLACE PROCEDURE add_product(
  p_category VARCHAR,
  p_name VARCHAR,
  p_quantity INT,
  p_price NUMERIC(10,2)
)
LANGUAGE plpgsql
AS $$
DECLARE
  new_id INT;
BEGIN
  INSERT INTO product (category, name, quantity, price)
  VALUES (p_category, p_name, p_quantity, p_price)
  RETURNING id INTO new_id;

  RAISE NOTICE 'Продукт с id % добавлен.', new_id;
END;
$$;


CREATE OR REPLACE PROCEDURE delete_product_by_id(product_id INT)
LANGUAGE plpgsql
AS $$
BEGIN
  DELETE FROM product WHERE id = product_id;

  IF NOT FOUND THEN
    RAISE NOTICE 'Продукт с id % не найден.', product_id;
  ELSE
    RAISE NOTICE 'Продукт с id % удалён.', product_id;
  END IF;
END;
$$;

