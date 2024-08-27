import { ref, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
import getFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
import getTickets from "~/graphql/queries/tickets/getTickets.gql";

export default function useFetchData(userId) {
  const bookmarkedEvents = ref([]);
  const followingEvents = ref([]);
  const tickets = ref([]);

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
  } = useQuery(
    getBookmarkedEvents,
    { id: userId },
    { pollInterval: pollingInterval }
  );
  watch(bookmarkedEventsResult, (newResult) => {
    if (newResult?.bookmarkedEvents) {
      bookmarkedEvents.value = newResult.bookmarkedEvents;
    }
  });

  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
  } = useQuery(
    getFollowingEvents,
    { id: userId },
    { pollInterval: pollingInterval }
  );
  watch(followingEventsResult, (newResult) => {
    if (newResult?.followingEvents) {
      followingEvents.value = newResult.followingEvents;
    }
  });

  const {
    result: ticketsResult,
    loading: loadingTickets,
    error: ticketsError,
  } = useQuery(getTickets, { id: userId }, { pollInterval: pollingInterval });
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
  };
}
