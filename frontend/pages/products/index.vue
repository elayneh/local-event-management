<template>
  <div>
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6">
      <CustomCard
        v-for="product in products"
        :key="product.id"
        :title="product.name"
        :image="`data:image/jpeg;base64,${product.image}`"
      >
        <template #body>
          <p class="text-gray-500 text-sm">Category: {{ product.category }}</p>
          <p class="text-gray-700">{{ product.description }}</p>
          <p class="text-lg font-semibold text-blue-500 mt-2">
            Price: ${{ product.price }}
          </p>
        </template>
      </CustomCard>
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 p-6 w-50 w-100"
      >
        <NuxtLink to="/products/add">
          <img src="./../../assets/images/add.png" width="50px" />
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import CustomCard from "~/components/CustomCard.vue";
import { ref, onMounted, watch } from "vue";
import { useQuery } from "@vue/apollo-composable";
import getProducts from "~/graphql/queries/products/getProducts.gql";

const products = ref([]);
const loading = ref(false);
const error = ref(null);

definePageMeta({ layout: "authenticated" });

onMounted(() => {
  fetchProducts();
});

function fetchProducts() {
  const {
    result,
    loading: queryLoading,
    error: queryError,
  } = useQuery(getProducts);

  watch(result, (newVal) => {
    result.value = newVal;
  });
  watch(queryLoading, (newVal) => {
    loading.value = newVal;
  });

  watch(queryError, (newVal) => {
    error.value = newVal;
  });

  watch(result, (newResult) => {
    if (newResult?.products) {
      products.value = newResult.products;
    }
  });
}
</script>
