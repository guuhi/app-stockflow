INSERT INTO products (sku, name, price, stock, is_active, created_at) 
VALUES
    ('SKU-001', 'Kopi Arabika Gayo 250g', 85000, 42, TRUE, now() - INTERVAL '30 days'),
    ('SKU-002', 'Teh Hijau Melati Premium 100g', 45000, 110, TRUE, now() - INTERVAL '28 days'),
    ('SKU-003', 'Gula Aren Organik Bubuk 500g', 38000, 75, TRUE, now() - INTERVAL '25 days'),
    ('SKU-004', 'Tumbler Stainless 500ml - Hitam', 149000, 18, TRUE, now() - INTERVAL '22 days'),
    ('SKU-005', 'French Press Kaca 600ml', 185000, 12, TRUE, now() - INTERVAL '20 days'),
    ('SKU-006', 'V60 Coffee Dripper Ceramic', 120000, 0, TRUE, now() - INTERVAL '18 days'),
    ('SKU-007', 'Kertas Filter V60 (Isi 40)', 35000, 200, TRUE, now() - INTERVAL '15 days'),
    ('SKU-008', 'Manual Coffee Grinder Conical Burr', 325000, 9, TRUE, now() - INTERVAL '14 days'),
    ('SKU-009', 'Timbangan Kopi Digital dengan Timer', 210000, 15, TRUE, now() - INTERVAL '12 days'),
    ('SKU-010', 'Gooseneck Kettle Stainless 1L', 275000, 6, TRUE, now() - INTERVAL '10 days'),
    ('SKU-011', 'Cangkir Keramik Handmade 200ml', 65000, 30, TRUE, now() - INTERVAL '7 days'),
    ('SKU-012', 'Madu Hutan Sumbawa 350ml', 115000, 24, TRUE, now() - INTERVAL '5 days'),
    ('SKU-013', 'Biskuit Speculoos Caramel 250g', 42000, 50, TRUE, now() - INTERVAL '3 days'),
    ('SKU-014', 'Paket Hampers Kopi Nusantara', 450000, 0, FALSE, now() - INTERVAL '45 days'),
    ('SKU-015', 'Voucher Belanja Digital Rp 50.000', 50000, 999, TRUE, now() - INTERVAL '1 day')
ON CONFLICT (sku) DO NOTHING;