import { ref } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import GetFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";
import getUserEvents from "~/graphql/queries/events/getUserEvents.gql";
import getSingleUser from "~/graphql/queries/users/getSingleUser.gql";

const bookmarkedEvents = ref([]);
const followingEvents = ref([]);
const tickets = ref([]);
const userEvents = ref([]);
const user = ref({});

export default function useUserFetchData(user_id) {
  const {
    loading: userEventsLoading,
    error: userEventsError,
    onResult: onUserEventsResult,
  } = useQuery(
    getUserEvents,
    { id: user_id },
    {
      fetchPolicy: "no-cache",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "user",
          "x-hasura-is-email-verified": true,
        },
      },
    }
  );

  onUserEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      userEvents.value = newResult;
    } else {
      console.error("Error: events not found");
    }
  });

  const {
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
    onResult: onBookmarkedEventsResult,
  } = useQuery(
    getBookmarkedEvents,
    { id: user_id },
    {
      fetchPolicy: "no-cache",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "user",
          "x-hasura-is-email-verified": true,
        },
      },
    }
  );

  onBookmarkedEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      bookmarkedEvents.value = newResult;
    } else {
      console.error("Error: bookmarked events not found");
    }
  });

  const {
    loading: loadingFollowingEvents,
    error: followingEventsError,
    onResult: onFollowingEventsResult,
  } = useQuery(
    GetFollowingEvents,
    { id: user_id },
    {
      fetchPolicy: "no-cache",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "user",
          "x-hasura-is-email-verified": true,
        },
      },
    }
  );

  onFollowingEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      followingEvents.value = newResult;
    } else {
      console.error("Error: following events not found");
    }
  });

  const {
    loading: loadingTickets,
    error: ticketsError,
    onResult: onTicketsResult,
  } = useQuery(getTickets, { id: user_id });

  onTicketsResult((response) => {
    const newResult = response.data?.tickets;
    if (newResult) {
      tickets.value = newResult;
    } else {
      console.error("Error: tickets not found");
    }
  });

  const {
    loading: userLoading,
    error: userError,
    onResult: onUserResult,
  } = useQuery(getSingleUser, { id: user_id });

  onUserResult((response) => {
    const newResult = response.data?.users_by_pk;
    if (newResult) {
      user.value = newResult;
    } else {
      console.error("Error: user not found");
    }
  });

  return {
    user,
    followingEvents,
    tickets,
    bookmarkedEvents,
    userEvents,
    userLoading,
    userError,
    loadingTickets,
    ticketsError,
    loadingBookmarkedEvents,
    bookmarkedEventsError,
    loadingFollowingEvents,
    followingEventsError,
    userEventsLoading,
    userEventsError,
  };
}