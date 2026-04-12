INSERT INTO
  users (id, email, password_hash, created_at)
VALUES
  (
    '00000000-0000-0000-0000-000000000001',
    'test@example.com',
    'dummy-password-hash',
    NOW()
  )
ON CONFLICT (id) DO NOTHING;
