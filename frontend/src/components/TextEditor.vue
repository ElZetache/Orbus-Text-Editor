<script setup lang="ts">
	import { ref, watch, onMounted } from "vue";

	const lines = ref("");
	const numbers = ref<number[]>([1]);
	const cursorLine = ref<number>(0);
	const textarea = ref<HTMLTextAreaElement | null>(null);

	const countNewLines = (str: string): number => {
		return (str.match(/\n/g) || []).length;
	};

	const getCursorLine = (textarea: HTMLTextAreaElement): number => {
		const cursorPosition = textarea.selectionStart;
		const textUntilCursor = textarea.value.substr(0, cursorPosition);
		return textUntilCursor.split("\n").length;
	};

	watch(lines, (newVal, oldVal) => {
		const oldLinesCount = countNewLines(oldVal);
		const newLinesCount = countNewLines(newVal);

		const difference = newLinesCount - oldLinesCount;
		if (difference > 0) {
			for (let i = 0; i < difference; i++) {
				numbers.value.push(numbers.value.length + 1);
			}
		} else if (difference < 0) {
			for (let i = 0; i < Math.abs(difference); i++) {
				numbers.value.pop();
			}
		}

		adjustTextareaWidth();
	});

	const updateCursorLine = (event: Event) => {
		const textarea = event.target as HTMLTextAreaElement;
		cursorLine.value = getCursorLine(textarea);
	};

	const adjustTextareaWidth = () => {
		if (textarea.value) {
			textarea.value.style.width = "auto";
			textarea.value.style.width = textarea.value.scrollWidth + "px";
		}
	};

	onMounted(() => {
		adjustTextareaWidth();
	});
</script>

<template>
	<main class="flex">
		<section class="w-screen flex">
			<div class="block text-right min-h-[93vh] ml-1">
				<p
					class="text-[#5f5f5f]"
					v-for="(number, index) in numbers"
					:key="index"
					:class="{ 'text-white': cursorLine === number }"
				>
					{{ number }}
				</p>
			</div>
			<textarea
				v-model="lines"
				ref="textarea"
				class="border-none focus:outline-none focus:ring-0 min-w-[97vw] min-h-[93vh] p-1 resize-none bg-[#0f0f0f] "
				name="textEditor"
				id="textEditor"
				:placeholder="lines != '' ? '' : 'Type here...'"
				@input="updateCursorLine"
				@click="updateCursorLine"
				@keyup="updateCursorLine"
				>{{ lines }}</textarea
			>
		</section>
	</main>
</template>

<style scoped>
	textarea {
		white-space: nowrap;
	}
</style>
