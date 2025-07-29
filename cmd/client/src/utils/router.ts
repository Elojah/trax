import router from "@/router";

/*
  * Utility function to navigate back in the router history.
 */
export function back() {
  router.go(-1);
};
