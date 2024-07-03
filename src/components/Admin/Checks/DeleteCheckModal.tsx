import { ConfirmModal } from "../..";

type props = {
  check: string;
  open: boolean;
  setOpen: (open: boolean) => void;
  handleDelete: () => void;
};

export default function DeleteCheckModal({
  check,
  open,
  setOpen,
  handleDelete,
}: props) {
  return (
    <ConfirmModal
      title='Delete Check'
      subtitle={
        <>
          To confirm deletion of check, type the name (<b>{check}</b>) of the
          check below.
        </>
      }
      buttonText='Delete Check'
      value={check}
      open={open}
      setOpen={setOpen}
      onConfirm={handleDelete}
      label='Check Name'
    />
  );
}
