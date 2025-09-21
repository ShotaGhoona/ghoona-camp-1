export const APP_CONFIG = {
  APP_NAME: 'Ghoona Camp',
  APP_DESCRIPTION: '朝活コミュニティアプリ',
  APP_URL: process.env.NEXT_PUBLIC_APP_URL || 'http://localhost:3000',
} as const;

export const API_CONFIG = {
  BASE_URL: process.env.NEXT_PUBLIC_API_URL || '/api',
  TIMEOUT: 10000,
} as const;

export const ROUTES = {
  HOME: '/',
  DASHBOARD: '/dashboard',
  PROFILE: '/profile',
  GOALS: '/goals',
  EVENTS: '/events',
  RANKING: '/ranking',
  SIGN_IN: '/sign-in',
  SIGN_UP: '/sign-up',
} as const;

export const DISCORD_CONFIG = {
  MORNING_HOURS: {
    START: 6,
    END: 6.5, // 6:30
  },
} as const;
