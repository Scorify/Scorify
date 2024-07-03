import { ConfirmModal } from "../..";

type props = {
  inject: string;
  open: boolean;
  setOpen: (open: boolean) => void;
  handleDelete: () => void;
};

export default function DeleteInjectModal({
  inject,
  open,
  setOpen,
  handleDelete,
}: props) {
  return (
    <ConfirmModal
      title='Delete Inject'
      subtitle={
        <>
          To confirm deletion of inject, type the name (<b>{inject}</b>) of the
          inject below.
        </>
      }
      buttonText='Delete Inject'
      value={inject}
      open={open}
      setOpen={setOpen}
      onConfirm={handleDelete}
      label='Inject Name'
    />
  );
}
