/**
 * Query keys for event-related queries
 */
export const eventQueryKeys = {
  all: ['event'] as const,
  
  // Event queries
  events: () => [...eventQueryKeys.all, 'events'] as const,
  eventsList: (filters?: Record<string, any>) => 
    [...eventQueryKeys.events(), 'list', filters] as const,
  eventsDetail: (id: string) => 
    [...eventQueryKeys.events(), 'detail', id] as const,
  
  // Participants queries
  participants: () => [...eventQueryKeys.all, 'participants'] as const,
  participantsList: (eventId: string, filters?: Record<string, any>) => 
    [...eventQueryKeys.participants(), 'list', eventId, filters] as const,
  participantsDetail: (eventId: string, participantId: string) => 
    [...eventQueryKeys.participants(), 'detail', eventId, participantId] as const,
} as const;