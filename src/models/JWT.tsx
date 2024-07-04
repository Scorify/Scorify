export type JWT =
  | {
      become: string | undefined;
      id: string;
      exp: number;
    }
  | undefined;
