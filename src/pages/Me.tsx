import { useMeQuery } from "../graph";

export default function Me() {
  const { data, loading, error } = useMeQuery();

  return (
    <div>
      {loading && <p>Loading...</p>}
      {error && <p>Error: {JSON.stringify(error)}</p>}
      {data && <p>{JSON.stringify(data.me)}</p>}
    </div>
  );
}
