import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export function useURLParam<U>(
  key: string,
  // Convert the value to a string
  encode: (value: U) => string,
  // Convert the string to a value
  decode: (s: string) => U
): {
  parameter: U | undefined;
  setParameter: React.Dispatch<React.SetStateAction<U | undefined>>;
  deleteParameter: () => void;
} {
  const navigate = useNavigate();
  const urlSearchParams = new URLSearchParams(location.search);

  const setUrlParam = (value: string) => {
    urlSearchParams.set(key, value);
    navigate(`?${urlSearchParams.toString()}`);
  };
  const getUrlParam = () => urlSearchParams.get(key);
  const deleteParameter = () => urlSearchParams.delete(key);

  const raw = getUrlParam();
  const [parameter, setParameter] = useState<U | undefined>(
    raw === null ? undefined : decode(raw)
  );
  useEffect(() => {
    if (parameter === undefined || parameter === "") {
      deleteParameter();
    } else {
      setUrlParam(encode(parameter));
    }
  }, [parameter]);

  return { parameter, setParameter, deleteParameter };
}
