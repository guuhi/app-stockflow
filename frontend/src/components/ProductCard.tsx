interface ProductCardProps {
  sku: string;
  name: string;
  price: number;
  stock: number;
  isActive: boolean;
  createdAt: string;
  lastMovementAt: string;
}

function ProductCard({
  sku,
  name,
  price,
  stock,
  isActive,
  createdAt,
  lastMovementAt,
}: ProductCardProps) {
  return (
    <div className="card flex flex-col gap-2">
      <div className="flex justify-between">
        <h2 className="font-bold">
          {sku} - {name}
        </h2>
        <p
          className={`rounded-xl px-3 py-1 text-sm ${
            isActive ? 'bg-green-200 ' : 'bg-red-200'
          }`}
        >
          {formatStatus(isActive)}
        </p>
      </div>

      <p>Price: {priceFormatter(price)}</p>
      <p>Stock: {stock}</p>

      <div className="flex justify-between text-slate-500">
        <p>CreatedAt: {formatDate(createdAt)}</p>
        <p>LastMovementAt: {formatDate(lastMovementAt)}</p>
      </div>
    </div>
  );
}

export default ProductCard;

function priceFormatter(price: number): string {
  const formatted = price.toLocaleString('id-ID', {
    style: 'currency',
    currency: 'IDR',
  });
  return formatted;
}

function formatStatus(isActive: boolean): string {
  if (isActive === true) {
    return 'Active';
  } else {
    return 'Inactive';
  }
}

function formatDate(date: string): string {
  const newDate = new Date(date);
  return newDate.toLocaleDateString();
}
