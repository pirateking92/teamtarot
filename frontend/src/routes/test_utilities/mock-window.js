import { vi } from 'vitest';

export const mockWindowLanding = {
    location: {
        href: '/',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};

export const mockWindowCassandra = {
    location: {
        href: '/cassandra?name=Paul',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};

export const mockWindowCardDisplay = {
    location: {
        href: '/cards',
        assign: vi.fn(),
        replace: vi.fn(),
    },
};