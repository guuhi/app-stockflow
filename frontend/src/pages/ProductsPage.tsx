import useSWR from 'swr';
import SearchBar from '../components/SearchBar';
import { useSearchParams } from 'react-router';
import { type ListProductsResponse, type Product } from '../types';
import ProductCard from '../components/ProductCard';
import PaginationBar from '../components/PaginationBar';

export const API_URL: string =
  import.meta.env.VITE_API_URL ?? 'http://localhost:8081';

const fetcher = (path: string) =>
  fetch(`${API_URL}${path}`).then((res) => {
    return res.json() as Promise<ListProductsResponse>;
  });

function ProductsPage() {
  const [searchParams, setSearchParams] = useSearchParams();
  const queryString = searchParams.toString();

  const key = queryString ? `/api/products?${queryString}` : '/api/products';

  const { data, error, isLoading } = useSWR<ListProductsResponse>(key, fetcher);

  const page: number = data ? data.page : 1;
  const total: number = data ? data.total : 1;
  const products: Product[] = data ? data.data : [];

  const productCards = products.map((p) => (
    <ProductCard
      key={p.id}
      sku={p.sku}
      name={p.name}
      stock={p.stock}
      price={p.price}
      isActive={p.is_active}
      createdAt={p.created_at}
      lastMovementAt={p.last_movement_at}
    />
  ));

  return (
    <div className="px-8 py-5 flex flex-col gap-4">
      <SearchBar onSearchQueryChange={setSearchParams} />
      {isLoading && <p data-state="loading">Loading products data...</p>}
      {!isLoading && (error || !data) && (
        <p className="text-red-600" data-state="error">
          {error?.message ?? 'Failed to load products.'}
        </p>
      )}
      {!isLoading && !error && products.length === 0 && (
        <p data-state="loading">No products...</p>
      )}
      {!isLoading && !error && (
        <div>
          <div className="grid grid-cols-2 gap-4">{productCards}</div>
          <div className="py-4">
            <PaginationBar
              currentPage={page}
              totalPages={total}
              onPageChange={setSearchParams}
            />
          </div>
        </div>
      )}
    </div>
  );
}

export default ProductsPage;
