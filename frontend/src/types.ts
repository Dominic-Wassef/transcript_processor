// Define the Utterance type for your transcript data
export type Utterance = {
    groupId: number;
    endTime: string;
    meetingId: string;
    offsetSeconds: string;
    speaker: string;
    startTime: string;
    text: string;
    timestampMs: number;
    timestamp: string;
  };
  