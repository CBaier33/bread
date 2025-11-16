import React, { useState } from "react";
import { Dialog, Popover, Text, TextField, Flex, Button, IconButton} from "@radix-ui/themes";
import { PlusIcon } from "@radix-ui/react-icons";

interface NewGroupPromptProps {
  onSave: (name: string, description: string) => Promise<void> | void;
}

const NewGroupPrompt: React.FC<NewGroupPromptProps> = ({ onSave }) => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");


  const handleCreateGroup = async () => {
    await onSave(name, description);
  };

  return (
      <Dialog.Root>
        <Dialog.Trigger>
          <IconButton size="1" color="gray" radius="full" variant="solid">
            <PlusIcon />
          </IconButton>
        </Dialog.Trigger>
        <Flex direction="column" width="32rem" asChild>
        <Dialog.Content>
          <Dialog.Title>New Group</Dialog.Title>
          <Flex gap="2" direction="column">
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Name
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  defaultValue=""
                  value={name}
                  placeholder=""
                  onChange={(e) => setName(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
            <label>
              <Text as="div" size="2" mb="1" weight="bold">
                Description
              </Text>
              <Flex width="20rem" asChild>
                <TextField.Root
                  defaultValue=""
                  value={description}
                  placeholder=""
                  onChange={(e) => setDescription(e.target.value)}
                  size="2"
                />
              </Flex>
            </label>
          </Flex>
            <Dialog.Close>
              <Flex gap="3" mt="4" justify="end">
                  <Button variant="soft" color="gray">
                    Cancel
                  </Button>
                  <Button onClick={handleCreateGroup}>
                    Save
                  </Button>
              </Flex>
            </Dialog.Close>
        </Dialog.Content>
        </Flex>

      </Dialog.Root>
  );
};

export default NewGroupPrompt;
