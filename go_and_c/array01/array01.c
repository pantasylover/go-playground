        if (!ret->entries[i])
            continue;

        rLen++;

        memcpy(ret->entries[i], &people[i], sizeof(Person));
    }

    ret->length = rLen;

    return ret;
}

void FreePeople(People* p) {
    if (p) {
        for (int i = 0; i < p->length; i++) {
            Person *t = p->entries[i];

            if (t)
                free(t);
        }

        free(p);
    }
}
