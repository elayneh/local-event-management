export default (filter, order, limit, offset, getTodos) => {
  const { onResult, loading, refetch, onError } = useQuery(
    getTodos,
    () => ({
      offset: offset.value,
      limit: limit.value,
      order: order.value,
      filter: filter.value,
    }),
    () => ({
      fetchPolicy: "network-only",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "user",
          "x-hasura-is-email-verified": true,
        },
      },
    })
  );

  return {
    onResult,
    loading,
    refetch,
    onError,
  };
};
