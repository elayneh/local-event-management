import { useQuery } from "@vue/apollo-composable";

export default (id, getData) => {
  const { onResult, loading, refetch, onError } = useQuery(
    getData,
    () => ({
      id,
    }),
    {
      fetchPolicy: "network-only",
    }
  );

  return {
    onResult,
    loading,
    refetch,
    onError,
  };
};
