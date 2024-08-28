import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import getFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";
import getUserEvents from "~/graphql/queries/events/getUserEvents.gql";

export default function useUserFetchData(user_id) {
  const bookmarkedEvents = ref([]);
  const followingEvents = ref([]);
  const tickets = ref([]);
  const categories = ref([]);
  const tags = ref([]);
  const userEvents = ref([]);

  const pollingInterval = 10000;

  const {
    result: userEventsResult,
    loading: userEventsLoading,
    error: userEventsError,
  } = useQuery(getUserEvents, { id: user_id });

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
  } = useQuery(
    getBookmarkedEvents,
    { id: user_id },
    { pollInterval: pollingInterval }
  );
  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
  } = useQuery(
    getFollowingEvents,
    { id: user_id },
    { pollInterval: pollingInterval }
  );
  const {
    result: ticketsResult,
    loading: loadingTickets,
    error: ticketsError,
  } = useQuery(getTickets, { id: user_id }, { pollInterval: pollingInterval });

  watch(userEventsResult, (newResult) => {
    if (newResult?.events) {
      userEvents.value = newResult.events;
    }
  });

  watch(bookmarkedEventsResult, (newResult) => {
    if (newResult?.events) {
      bookmarkedEvents.value = newResult.events;
    }
  });

  watch(followingEventsResult, (newResult) => {
    if (newResult?.events) {
      followingEvents.value = newResult.events;
    }
  });

  watch(ticketsResult, (newResult) => {
    if (newResult?.tickets) {
      tickets.value = newResult.tickets;
    }
  });

  return {
    followingEvents,
    tickets,
    bookmarkedEvents,
    loadingBookmarkedEvents,
    bookmarkedEventsError,
    userEvents,
    userEventsLoading,
    userEventsError,
    categories,
    tags,
  };
}
