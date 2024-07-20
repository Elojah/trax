import { grpccodes } from "@/utils/errors"
import { defineStore } from "pinia"
import { ref } from "vue"

export const useErrorsStore = defineStore('errors', () => {
  const message = ref<string>('')

  const success = ref<boolean>(false)
  const error = ref<boolean>(false)

  const showGRPC = (e: any) => {
    switch (e.code) {
      case grpccodes.OK:
        message.value = 'Success.';
        success.value = true;
        break;
      case grpccodes.INTERNAL:
        message.value = 'Internal error occurred, please retry later.';
        error.value = true;
        break;
      case grpccodes.UNAUTHENTICATED:
        message.value = 'Authentication required to perform this action.';
        error.value = true;
        break;
      case grpccodes.PERMISSION_DENIED:
        message.value = 'Missing permission to perform this action.';
        error.value = true;
        break;
      case grpccodes.INVALID_ARGUMENT:
        message.value = 'Invalid options provided, please change your request.';
        error.value = true;
        break;
      case grpccodes.NOT_FOUND:
        message.value = 'Resource not found.';
        error.value = true;
        break;
      case grpccodes.ALREADY_EXISTS:
        message.value = 'Resource already exists.';
        error.value = true;
        break;

      case grpccodes.CANCELLED:
        message.value = 'Operation cancelled.';
        error.value = true;
        break;
      case grpccodes.UNKNOWN:
        message.value = `Unknown error: ${e.message}.`;
        error.value = true;
        break;
      case grpccodes.DEADLINE_EXCEEDED:
        message.value = 'Operation taking too long to complete.';
        error.value = true;
        break;
      case grpccodes.RESOURCE_EXHAUSTED:
        message.value = 'Resource no more available.';
        error.value = true;
        break;
      case grpccodes.FAILED_PRECONDITION:
        message.value = 'Missing required condition.';
        error.value = true;
        break;
      case grpccodes.ABORTED:
        message.value = 'Operation aborted.';
        error.value = true;
        break;
      case grpccodes.OUT_OF_RANGE:
        message.value = 'Resource out of range.';
        error.value = true;
        break;
      case grpccodes.UNIMPLEMENTED:
        message.value = `Not implemented yet.Expected version: ${e.message}.`;
        error.value = true;
        break;
      case grpccodes.UNAVAILABLE:
        message.value = 'Resource unavailable.';
        error.value = true;
        break;
      case grpccodes.DATA_LOSS:
        message.value = 'Corrupted data.';
        error.value = true;
        break;
      case grpccodes.DO_NOT_USE:
        message.value = 'Not supported.';
        error.value = true;
        break

      default:
        message.value = `Unknown error with unknown code: ${e.message}.`;
        error.value = true;
        break;
    }
  }

  return {
    message,
    success,
    error,
    showGRPC,
  }
})
