// import { ref, watch, watchEffect } from "vue";
// import { useQuery } from "@vue/apollo-composable";
// import getBookmarkedEvents from "~/graphql/queries/bookmarks/getBookmarkedEvents.gql";
// import GetFollowingEvents from "~/graphql/queries/following/getFollowingEvents.gql";
// import getTickets from "~/graphql/queries/tickets/getTickets.gql";
// import getUserEvents from "~/graphql/queries/events/getUserEvents.gql";
// import getSingleUser from "~/graphql/queries/users/getSingleUser.gql";

// const bookmarkedEvents = ref([]);
// const followingEvents = ref([]);
// const tickets = ref([]);
// const userEvents = ref([]);
// const user = ref({});

// export default function useUserFetchData(user_id) {
//   const {
//     result: userEventsResult,
//     loading: userEventsLoading,
//     error: userEventsError,
//   } = useQuery(getUserEvents, { id: user_id });

//   watch(userEventsResult, (newResult) => {
//     if (newResult?.events) {
//       userEvents.value = newResult.events;
//     }
//   });

//   const {
//     result: bookmarkedEventsResult,
//     loading: loadingBookmarkedEvents,
//     error: bookmarkedEventsError,
//   } = useQuery(getBookmarkedEvents, { id: user_id });

//   watch(bookmarkedEventsResult, (newResult) => {
//     if (newResult?.events) {
//       bookmarkedEvents.value = newResult.events;
//     }
//   });

//   const {
//     result: followingEventsResult,
//     loading: loadingFollowingEvents,
//     error: followingEventsError,
//   } = useQuery(GetFollowingEvents, { id: user_id });

//   watch(followingEventsResult, (newResult) => {
//     if (newResult?.events) {
//       followingEvents.value = newResult.events;
//     }
//   });

//   const {
//     result: ticketsResult,
//     loading: loadingTickets,
//     error: ticketsError,
//     onResult: onTicketsResult,
//   } = useQuery(getTickets, { id: user_id });

//   watch(ticketsResult, (newResult) => {
//     if (newResult?.tickets) {
//       tickets.value = newResult.tickets;
//     }
//   });

//   const { result, loading, error, onResult } = useQuery(getSingleUser, {
//     id: user_id,
//   });

//   onResult((response) => {
//     const newResult = response.data?.users_by_pk;
//     if (newResult) {
//       user.value = newResult;
//     } else {
//       console.error("Error: users_by_pk not found");
//     }
//   });

//   return {
//     user,
//     followingEvents,
//     tickets,
//     bookmarkedEvents,
//     userEvents,
//   };
// }

//////////

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
    result: userEventsResult,
    loading: userEventsLoading,
    error: userEventsError,
    onResult: onUserEventsResult,
  } = useQuery(getUserEvents, { id: user_id });

  onUserEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      userEvents.value = newResult;
    } else {
      console.error("Error: events not found");
    }
  });

  const {
    result: bookmarkedEventsResult,
    loading: loadingBookmarkedEvents,
    error: bookmarkedEventsError,
    onResult: onBookmarkedEventsResult,
  } = useQuery(getBookmarkedEvents, { id: user_id });

  onBookmarkedEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      bookmarkedEvents.value = newResult;
    } else {
      console.error("Error: bookmarked events not found");
    }
  });

  const {
    result: followingEventsResult,
    loading: loadingFollowingEvents,
    error: followingEventsError,
    onResult: onFollowingEventsResult,
  } = useQuery(GetFollowingEvents, { id: user_id });

  onFollowingEventsResult((response) => {
    const newResult = response.data?.events;
    if (newResult) {
      followingEvents.value = newResult;
    } else {
      console.error("Error: following events not found");
    }
  });

  const {
    result: ticketsResult,
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
    result: userResult,
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
