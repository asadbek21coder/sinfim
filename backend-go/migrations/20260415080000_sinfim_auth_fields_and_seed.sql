-- +goose Up
-- +goose StatementBegin
ALTER TABLE auth.users ADD COLUMN IF NOT EXISTS phone_number VARCHAR;
ALTER TABLE auth.users ADD COLUMN IF NOT EXISTS full_name VARCHAR;
ALTER TABLE auth.users ADD COLUMN IF NOT EXISTS must_change_password BOOLEAN NOT NULL DEFAULT FALSE;

CREATE UNIQUE INDEX IF NOT EXISTS idx_users_phone_number_unique
    ON auth.users (phone_number)
    WHERE phone_number IS NOT NULL;

INSERT INTO auth.roles (name)
VALUES ('PLATFORM_ADMIN'), ('OWNER'), ('TEACHER'), ('MENTOR'), ('STUDENT')
ON CONFLICT (name) DO NOTHING;

WITH platform_admin_role AS (
    SELECT id FROM auth.roles WHERE name = 'PLATFORM_ADMIN'
), permissions(permission) AS (
    VALUES
        ('auth:user:read'),
        ('auth:user:manage'),
        ('auth:role:read'),
        ('auth:role:manage'),
        ('auth:access:read'),
        ('auth:access:manage'),
        ('auth:session:read'),
        ('auth:session:manage'),
        ('taskmill:view'),
        ('taskmill:manage'),
        ('alert:view'),
        ('alert:manage'),
        ('audit:action-log:read'),
        ('audit:status-change-log:read')
)
INSERT INTO auth.role_permissions (role_id, permission)
SELECT platform_admin_role.id, permissions.permission
FROM platform_admin_role, permissions
WHERE NOT EXISTS (
    SELECT 1
    FROM auth.role_permissions rp
    WHERE rp.role_id = platform_admin_role.id
      AND rp.permission = permissions.permission
);

INSERT INTO auth.users (id, username, phone_number, full_name, password_hash, is_active, must_change_password)
VALUES (
    '00000000-0000-4000-8000-000000000001',
    '+998900000001',
    '+998900000001',
    'Sinfim Platform Admin',
    '$2a$10$WLl.8IkkWrH4U0cEh/Z70usRQTMIJL6CgQQ88nIpKrSOFDhoSkCqC',
    TRUE,
    FALSE
)
ON CONFLICT (id) DO UPDATE SET
    username = EXCLUDED.username,
    phone_number = EXCLUDED.phone_number,
    full_name = EXCLUDED.full_name,
    is_active = TRUE,
    updated_at = CURRENT_TIMESTAMP;

WITH platform_admin_role AS (
    SELECT id FROM auth.roles WHERE name = 'PLATFORM_ADMIN'
)
INSERT INTO auth.user_roles (user_id, role_id)
SELECT '00000000-0000-4000-8000-000000000001', platform_admin_role.id
FROM platform_admin_role
WHERE NOT EXISTS (
    SELECT 1
    FROM auth.user_roles ur
    WHERE ur.user_id = '00000000-0000-4000-8000-000000000001'
      AND ur.role_id = platform_admin_role.id
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM auth.user_roles WHERE user_id = '00000000-0000-4000-8000-000000000001';
DELETE FROM auth.users WHERE id = '00000000-0000-4000-8000-000000000001';
DELETE FROM auth.role_permissions
WHERE role_id IN (SELECT id FROM auth.roles WHERE name = 'PLATFORM_ADMIN')
  AND permission IN (
      'auth:user:read', 'auth:user:manage', 'auth:role:read', 'auth:role:manage',
      'auth:access:read', 'auth:access:manage', 'auth:session:read', 'auth:session:manage',
      'taskmill:view', 'taskmill:manage', 'alert:view', 'alert:manage',
      'audit:action-log:read', 'audit:status-change-log:read'
  );
DELETE FROM auth.roles WHERE name IN ('PLATFORM_ADMIN', 'OWNER', 'TEACHER', 'MENTOR', 'STUDENT');
DROP INDEX IF EXISTS auth.idx_users_phone_number_unique;
ALTER TABLE auth.users DROP COLUMN IF EXISTS must_change_password;
ALTER TABLE auth.users DROP COLUMN IF EXISTS full_name;
ALTER TABLE auth.users DROP COLUMN IF EXISTS phone_number;
-- +goose StatementEnd
