<template>
  <section class="product-area shop-sidebar shop section">
    <div class="container">
      <div class="row">
        <LeftBar />
        <div class="col-lg-9 col-md-8 col-12">
          <div class="row">
            <div class="col-12">
								<!-- Shop Top -->
								<div class="shop-top">
									<div class="shop-shorter">
                    <div class="single-shorter">
											<label>Sort By :</label>
											<select >
												<option selected="selected">Name</option>
												<option>Price</option>
                      </select>
										</div>
										<div class="single-shorter">
											<label>Show :</label>
											<select>
												<option>Defaut</option>
												<option>ASC</option>
												<option>DESC</option>
											</select>
										</div>
									</div>
								</div>
								<!--/ End Shop Top -->
							</div>
          </div>
          <div class="row" v-if="!isLoading">
            <div
              class="col-lg-4 col-md-6 col-12"
              v-for="product in products"
              :key="product.ID"
            >
              <div class="single-product">
                <div class="product-img">
                  <router-link :to="'/product/' + product.ID">
                    <img
                      v-if="product.Url !== ''"
                      class="default-img"
                      :src="product.Url"
                      alt="#"
                    />
                  </router-link>
                  <div class="button-head">
                    <div class="product-action">
                      <a
                        data-toggle="modal"
                        data-target="#exampleModal"
                        title="Quick View"
                        href="#"
                        ><i class="ti-eye"></i><span>Quick Shop</span></a
                      >
                    </div>
                    <div class="product-action-2">
                      <a title="Add to cart" @click="addToCart(product)"
                        >Add to cart</a
                      >
                    </div>
                  </div>
                </div>
                <div class="product-content">
                  <h3>
                    <router-link :to="'/product/' + product.ID">
                      {{ product.Name }}
                    </router-link>
                  </h3>
                  <div class="product-price">
                    <span>${{ product.Price }}</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="col-lg-12 col-md-12 col-12">
               <Pagination
              :length="totalItems"
              :pageSize="limit"
              :pageIndex="pageIndex"
              @change="changePage"
            />
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import LeftBar from "./LeftBar.vue";
import { mapState, mapGetters, mapActions } from "vuex";
import Pagination from "@/components/Pagination.vue";

export default {
  name: "ProductsPage",

  components: {
    LeftBar,
    Pagination,
  },

  computed: {
    ...mapState("products", [
      "isLoading",
      "products",
      "totalItems",
      "pageIndex",
      "limit",
    ]),
    ...mapGetters("products", [
      "sortDropdownValue",
      "itemStartIndex",
      "itemEndIndex",
    ]),
  },

  created() {
    this.$store.dispatch("products/getProducts", {});
  },

  methods: {
    // currency,
    sortProducts(option) {
      const options = option.value.split("-");
      const sort = options[0],
        order = options[1];
      this.$store.dispatch("products/getProducts", { sort, order });
    },
    changePage(pageIndex) {
      this.$store.dispatch("products/getProducts", { pageIndex });
    },
    ...mapActions("products", ["getProducts"]),
  },
};
</script>