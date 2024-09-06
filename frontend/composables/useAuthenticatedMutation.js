export default function useAuthenticatedMutation(mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "no-cache",
    clientId: "authClient",
    context: {
      headers: {
        "x-hasura-role": "user",
        "x-hasura-is-email-verified": true,w
      },
    },
  }));

  return {
    mutate,
    onDone,
    loading,
    onError,
  };
}
