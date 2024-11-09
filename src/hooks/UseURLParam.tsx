import { useEffect, useState } from "react";

interface ConvertToString<U> {
  (value: U): string;
}

interface ConvertFromString<U> {
  (value: string): U;
}

export function useURLParam<U>(
  urlSearchParams: URLSearchParams,
  key: string,
  convertToString: ConvertToString<U>,
  convertFromString: ConvertFromString<U>
): [U | undefined, React.Dispatch<React.SetStateAction<U | undefined>>] {
  const raw = urlSearchParams.get(key);
  const [param, setParam] = useState<U | undefined>(
    raw === null ? undefined : convertFromString(raw)
  );
  useEffect(() => {
    if (param === undefined) {
      urlSearchParams.delete(key);
    } else {
      urlSearchParams.set(key, convertToString(param));
    }
  }, [param]);

  return [param, setParam];
}
