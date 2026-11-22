export interface Product {
  id: string;
  sku: string;
  name: string;
  price: number;
  stock: number;
  is_active: boolean;
  created_at: string;
  last_movement_at: string;
}

export interface ListProductsResponse {
  data: Product[];
  limit: number;
  page: number;
  total: number;
}
