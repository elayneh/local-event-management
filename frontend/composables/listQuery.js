const enable = ref(true);
const offset_ = ref(0);
const limit_ = ref(100);
const filter_ = ref({});

const order_ = ref([]);

export default function (
  query,
  {
    filter = filter_,
    order = order_,
    limit = limit_,
    offset = offset_,
    enabled = enable,
  }
) {
  const { onResult, onError, loading, refetch } = useQuery(
    query,
    () => ({
      limit: limit && limit?.value ? limit.value : undefined,
      filter: filter && filter.value ? filter.value : {},
      order: order && order?.value.length ? order.value : undefined,
      offset: offset && offset?.value ? offset.value : undefined,
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
}
