export type db_creator = {
  creator_id: number;
  creator_name: string;
};

export type db_sound_file = {
  sound_title: string;
  sound_created_time: Date;
  sound_updated_time: Date;
  sound_file_path: string;
  sound_file_type: string;
  sound_file_size: number;
  sound_text_result: string;
};
