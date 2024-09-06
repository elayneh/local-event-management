export default (id, getData) => {
	const { onResult, loading, refetch, onError } = useQuery(
    getData,
    () => ({
      id: id,
    }),
    () => ({
      fetchPolicy: "no-cache",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "user",
          "x-hasura-is-email-verified": true,
        },
      },
      enabled: enabled.value,
    })
  );

	return {
		onResult,
		loading,
		refetch,
		onError,
	};
};
