import { createBrowserRouter } from 'react-router';
import Layout from '../pages/Layout';
import ProductsPage from '../pages/ProductsPage';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [{ index: true, element: <ProductsPage /> }],
  },
]);
