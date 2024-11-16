import { useEffect, useState } from "react";

interface ConvertToString<U> {
  (value: U): string;
}

interface ConvertFromString<U> {
  (value: string): U;
}

export function useURLParam<U>(
  setUrlParam: (key: string, value: string) => void,
  getUrlParam: (key: string) => string | null,
  deleteUrlParam: (key: string) => void,
  key: string,
  convertToString: ConvertToString<U>,
  convertFromString: ConvertFromString<U>
): [U | undefined, React.Dispatch<React.SetStateAction<U | undefined>>] {
  const raw = getUrlParam(key);
  const [param, setParam] = useState<U | undefined>(
    raw === null ? undefined : convertFromString(raw)
  );
  useEffect(() => {
    console.log({ key, raw, param });
    if (param === undefined || param === "") {
      deleteUrlParam(key);
    } else {
      setUrlParam(key, convertToString(param));
    }
  }, [param]);

  return [param, setParam];
}
