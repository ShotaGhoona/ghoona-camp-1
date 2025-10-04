/**
 * Query keys for notification-related queries
 */
export const notificationQueryKeys = {
  all: ['notification'] as const,
  
  // Notifications queries
  notifications: () => [...notificationQueryKeys.all, 'notifications'] as const,
  notificationsList: (filters?: Record<string, unknown>) => 
    [...notificationQueryKeys.notifications(), 'list', filters] as const,
  notificationsDetail: (id: string) => 
    [...notificationQueryKeys.notifications(), 'detail', id] as const,
  notificationsUnread: () => 
    [...notificationQueryKeys.notifications(), 'unread'] as const,
  
  // Settings queries
  settings: () => [...notificationQueryKeys.all, 'settings'] as const,
  settingsProfile: () => 
    [...notificationQueryKeys.settings(), 'profile'] as const,
  settingsPreferences: () => 
    [...notificationQueryKeys.settings(), 'preferences'] as const,
} as const;